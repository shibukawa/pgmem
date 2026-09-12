package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_dim(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v10 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v9 & int32(2147483647)
			}
		} else {
			return v9 & int32(2147483647)
		}
	}
}
func F_cube_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v488 != v6 {
		goto L123
	} else {
		goto L124
	}
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
	v38 = v11 + v37
	v40 = v6 + v37
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
	v108 = v11 + v107
	v110 = v6 + v107
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
	v197 = v6 + int32(8)
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
	v312 = v11 + int32(8)
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
L123:
	;
	F_pfree(m, v6)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v492 != v11 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	F_pfree(m, v11)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	return base.B2i32(v487 == int32(0))
L130:
	;
	goto L129
}
func F_cube_inter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v117 int32
	_ = v117
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v127 float64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v138 float64
	_ = v138
	var v149 float64
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v160 float64
	_ = v160
	var v162 float64
	_ = v162
	var v166 float64
	_ = v166
	var v168 int32
	_ = v168
	var v171 float64
	_ = v171
	var v176 float64
	_ = v176
	var v183 float64
	_ = v183
	var v189 float64
	_ = v189
	var v190 float64
	_ = v190
	var v192 float64
	_ = v192
	var v198 int32
	_ = v198
	var v199 float64
	_ = v199
	var v200 int32
	_ = v200
	var v201 float64
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 float64
	_ = v219
	var v223 int32
	_ = v223
	var v230 float64
	_ = v230
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 float64
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 float64
	_ = v316
	var v318 int32
	_ = v318
	var v319 float64
	_ = v319
	var v322 float64
	_ = v322
	var v324 float64
	_ = v324
	var v326 float64
	_ = v326
	var v329 float64
	_ = v329
	var v330 float64
	_ = v330
	var v332 float64
	_ = v332
	var v335 float64
	_ = v335
	var v337 float64
	_ = v337
	var v340 float64
	_ = v340
	var v343 float64
	_ = v343
	var v344 float64
	_ = v344
	var v347 float64
	_ = v347
	var v348 int32
	_ = v348
	var v352 float64
	_ = v352
	var v358 int32
	_ = v358
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v425 int32
	_ = v425
	var v428 float64
	_ = v428
	var v433 float64
	_ = v433
	var v436 int32
	_ = v436
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v39 = int32(2147483647)
	v40 = v38 & v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v43 = v41 & v39
	v44 = base.B2i32(base.Ui32(v40) < base.Ui32(v43))
	if base.Ui32(v40) < base.Ui32(v43) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v45 = v36
	goto L6
L5:
	;
	v45 = v31
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v50 = v46<<(uint(int32(4))%32) | int32(8)
	v51 = F_palloc0(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v50 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v61 = v56&int32(-2147483648) | v46&int32(2147483647)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.Ui32(v40) < base.Ui32(v43) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui32(v67) < base.Ui32(v246) {
		goto L61
	} else {
		goto L62
	}
L9:
	;
	v64 = v31
	goto L11
L10:
	;
	v64 = v36
	goto L11
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v67 = v65 & int32(2147483647)
	if v67 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v246 = v63 & int32(2147483647)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v73 = v63 & int32(2147483647)
	v74 = int32(8)
	v75 = v51 + v74
	v77 = v64 + v74
	v79 = v45 + v74
	v80 = int32(0)
	v84 = v80
	goto L15
L15:
	;
	v113 = v84 << (uint(int32(3)) % 32)
	v114 = v79 + v113
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	v117 = base.B2i32(v63 < int32(0))
	if v63 < int32(0) {
		v125 = v115
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v246 = v73
	goto L8
L17:
	;
	v126 = v113 + v77
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v126)))
	v128 = int32(0)
	v129 = base.B2i32(v65 < v128)
	if v129 == v128 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v79+(v84+v63)<<(uint(int32(3))%32))))
	if base.F64_gt(v122, v115) != 0 {
		v125 = v115
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v125 = v122
	goto L17
L20:
	;
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v126)))
	if v129 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	v189 = v183
	v190 = v183
	goto L20
L22:
	;
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v79+v168<<(uint(int32(3))%32))))
	if base.F64_gt(v171, v176) == int32(0) {
		v189 = v171
		v190 = v176
		goto L20
	} else {
		goto L41
	}
L23:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v113+v75))) = v162
	v166 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	if v63 < int32(0) {
		v183 = v166
		goto L21
	} else {
		goto L40
	}
L24:
	;
	if base.B2i32(v80 <= v63) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v77+(v84+v65)<<(uint(int32(3))%32))))
	v137 = base.F64_lt(v127, v136)
	if v137 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if base.F64_lt(v127, v125) == int32(0) {
		v162 = v127
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v138 = v127
	goto L30
L29:
	;
	v138 = v136
	goto L30
L30:
	;
	if base.F64_gt(v125, v138) != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	if v137 != 0 {
		v162 = v127
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v162 = v136
	goto L23
L33:
	;
	goto L24
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v113+v75))) = v115
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	v189 = v149
	v190 = v149
	goto L20
L35:
	;
	goto L36
L36:
	;
	v151 = v84 + v63
	v154 = v79 + v151<<(uint(int32(3))%32)
	v155 = *(*float64)(unsafe.Add(mBase, uint32(v154)))
	if base.F64_lt(v115, v155) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v157 = v114
	goto L39
L38:
	;
	v157 = v154
	goto L39
L39:
	;
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v157)))
	*(*float64)(unsafe.Add(mBase, uint32(v113+v75))) = v158
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	v168 = v151
	v171 = v160
	goto L22
L40:
	;
	v168 = v84 + v63
	v171 = v166
	goto L22
L41:
	;
	v183 = v171
	goto L21
L42:
	;
	v230 = *(*float64)(unsafe.Add(mBase, uint32(v223)))
	*(*float64)(unsafe.Add(mBase, uint32(v75+(v84+v73)<<(uint(int32(3))%32)))) = v230
	v233 = v84 + int32(1)
	if v233 != v67 {
		v84 = v233
		goto L15
	} else {
		goto L60
	}
L43:
	;
	if v200 != 0 {
		v223 = v126
		goto L42
	} else {
		goto L59
	}
L44:
	;
	v218 = v79 + (v84+v63)<<(uint(int32(3))%32)
	v219 = *(*float64)(unsafe.Add(mBase, uint32(v218)))
	if base.F64_gt(v189, v219) != 0 {
		v223 = v114
		goto L42
	} else {
		goto L58
	}
L45:
	;
	v198 = v77 + (v84+v65)<<(uint(int32(3))%32)
	v199 = *(*float64)(unsafe.Add(mBase, uint32(v198)))
	v200 = base.F64_gt(v192, v199)
	if v200 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v207 = base.F64_gt(v192, v190)
	if v207 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v201 = v192
	goto L50
L49:
	;
	v201 = v199
	goto L50
L50:
	;
	if base.F64_lt(v190, v201) == int32(0) {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	if v117 == int32(0) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v223 = v114
	goto L42
L53:
	;
	v208 = v114
	goto L55
L54:
	;
	v208 = v126
	goto L55
L55:
	;
	if v63 < int32(0) {
		v223 = v208
		goto L42
	} else {
		goto L56
	}
L56:
	;
	if v207 == int32(0) {
		v223 = v208
		goto L42
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	v223 = v218
	goto L42
L59:
	;
	v223 = v198
	goto L42
L60:
	;
	goto L16
L61:
	;
	v265 = int32(8)
	v266 = v51 + v265
	v268 = v45 + v265
	v274 = v67
	goto L64
L62:
	;
	goto L63
L63:
	;
	v389 = int32(0)
	if v56 < v389 {
		goto L93
	} else {
		goto L94
	}
L64:
	;
	v301 = v274 << (uint(int32(3)) % 32)
	v302 = v268 + v301
	v303 = *(*float64)(unsafe.Add(mBase, uint32(v302)))
	if base.B2i32(v63 < int32(0)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L63
L66:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v266+v348<<(uint(int32(3))%32)))) = v352
	v358 = v274 + int32(1)
	if v358 != v246 {
		v274 = v358
		goto L64
	} else {
		goto L91
	}
L67:
	;
	v309 = int32(3)
	v311 = v268 + (v274+v63)<<(uint(v309)%32)
	v312 = v274 + v246
	v316 = *(*float64)(unsafe.Add(mBase, uint32(v268+v312<<(uint(v309)%32))))
	if base.F64_lt(v303, v316) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v337 = float64(0)
	if base.F64_lt(v303, v337) != 0 {
		goto L85
	} else {
		goto L86
	}
L70:
	;
	v318 = v302
	goto L72
L71:
	;
	v318 = v311
	goto L72
L72:
	;
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v318)))
	if base.F64_lt(v319, float64(0)) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v326 = float64(0)
	goto L75
L74:
	;
	v322 = *(*float64)(unsafe.Add(mBase, uint32(v311)))
	if base.F64_lt(v303, v322) != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v301+v266))) = v326
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v302)))
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v311)))
	if base.F64_gt(v329, v330) != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v324 = v303
	goto L78
L77:
	;
	v324 = v322
	goto L78
L78:
	;
	v326 = v324
	goto L75
L79:
	;
	v332 = v329
	goto L81
L80:
	;
	v332 = v330
	goto L81
L81:
	;
	if base.F64_gt(v332, float64(0)) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v335 = float64(0)
	goto L84
L83:
	;
	v335 = v332
	goto L84
L84:
	;
	v348 = v312
	v352 = v335
	goto L66
L85:
	;
	v340 = v337
	goto L87
L86:
	;
	v340 = v303
	goto L87
L87:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v301+v266))) = v340
	v343 = float64(0)
	v344 = *(*float64)(unsafe.Add(mBase, uint32(v302)))
	if base.F64_gt(v344, v343) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v347 = v343
	goto L90
L89:
	;
	v347 = v344
	goto L90
L90:
	;
	v348 = v274 + v246
	v352 = v347
	goto L66
L91:
	;
	goto L65
L92:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v40) < base.Ui32(v43) {
		goto L103
	} else {
		goto L104
	}
L93:
	;
	v470 = v46<<(uint(int32(3))%32) + int32(8)
	v471 = F_repalloc(m, v51, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L100
	}
L94:
	;
	if v61 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v395 = v51 + int32(8)
	v397 = v389
	goto L96
L96:
	;
	v425 = int32(3)
	v428 = *(*float64)(unsafe.Add(mBase, uint32(v395+v397<<(uint(v425)%32))))
	v433 = *(*float64)(unsafe.Add(mBase, uint32(v395+(v397+v61)<<(uint(v425)%32))))
	if base.F64_ne(v428, v433) != 0 {
		v486 = v51
		goto L92
	} else {
		goto L98
	}
L97:
	;
	goto L93
L98:
	;
	v436 = v397 + int32(1)
	if v436 != v61 {
		v397 = v436
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v470 << (uint(int32(2)) % 32)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v476 | int32(-2147483648)
	v486 = v471
	goto L92
L101:
	;
	return v486
L102:
	;
	F_pfree(m, v36)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L116
	}
L103:
	;
	if v509 != v31 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if v509 != v31 {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	F_pfree(m, v31)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v36 != v514 {
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	goto L101
L111:
	;
	F_pfree(m, v31)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v36 == v519 {
		goto L101
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	goto L102
L116:
	;
	goto L101
}
func F_cube_ll_coord(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v35 float64
	_ = v35
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v8 = float64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 <= int32(0) {
			v48 = v8
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v19 = v17 & int32(2147483647)
			if base.Ui32(v19) < base.Ui32(v14) {
				v48 = v8
			} else {
				v22 = v10 + int32(8)
				v24 = v14 - int32(1)
				v27 = v22 + v24<<(uint(int32(3))%32)
				if v17 < int32(0) {
					v41 = v27
				} else {
					v30 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
					v35 = *(*float64)(unsafe.Add(mBase, uint32(v22+(v24+v19)<<(uint(int32(3))%32))))
					if base.F64_lt(v30, v35) != 0 {
						v41 = v27
					} else {
						v41 = v22 + (v17+v24)<<(uint(int32(3))%32)
					}
				}
				v42 = *(*float64)(unsafe.Add(mBase, uint32(v41)))
				v48 = v42
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v49 != v10 {
			F_pfree(m, v10)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v53 = F_Float8GetDatum(m, v48)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					return v53
				}
			}
		} else {
			v53 = F_Float8GetDatum(m, v48)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return v53
			}
		}
	}
}
func F_cube_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 float64
	_ = v86
	var v91 float64
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_initStringInfo(m, v10)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoChar(m, v10, int32(40))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = v17 & int32(2147483647)
	if v24 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_appendStringInfoChar(m, v10, int32(41))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v27 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
	v28 = F_float8out_internal(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_appendStringInfoString(m, v10, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(1)
	if v24 == v32 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v38 = v32
	goto L10
L10:
	;
	F_appendStringInfoString(m, v10, int32(727489))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L5
L12:
	;
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v13+int32(8)+v38<<(uint(int32(3))%32))))
	v51 = F_float8out_internal(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_appendStringInfoString(m, v10, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v56 = v38 + int32(1)
	if v56 != v24 {
		v38 = v56
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v68 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v69 < v68 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v159 != v13 {
		goto L45
	} else {
		goto L46
	}
L18:
	;
	if v69 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v75 = v13 + int32(8)
	v77 = v68
	goto L20
L20:
	;
	v83 = int32(3)
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v75+v77<<(uint(v83)%32))))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v75+(v77+v69)<<(uint(v83)%32))))
	if base.F64_eq(v86, v91) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_appendStringInfoString(m, v10, int32(667418))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	v94 = v77 + int32(1)
	if v69 != v94 {
		v77 = v94
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L17
L26:
	;
	if v24 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_appendStringInfoChar(m, v10, int32(41))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L44
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v101 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v107 = v75
	goto L31
L30:
	;
	v107 = v75 + v101<<(uint(int32(3))%32)
	goto L31
L31:
	;
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v107)))
	v109 = F_float8out_internal(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_appendStringInfoString(m, v10, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v113 = int32(1)
	if v24 == v113 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v117 = v113
	goto L35
L35:
	;
	F_appendStringInfoString(m, v10, int32(727489))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L27
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v126 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v130 = v117
	goto L40
L39:
	;
	v130 = v126 + v117
	goto L40
L40:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v75+v130<<(uint(int32(3))%32))))
	v135 = F_float8out_internal(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_appendStringInfoString(m, v10, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v140 = v117 + int32(1)
	if v140 != v24 {
		v117 = v140
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	goto L17
L45:
	;
	F_pfree(m, v13)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	m.G0 = v10 + int32(16)
	return v163
L48:
	;
	goto L47
}
func F_cube_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pq_getmsgint(m, v12, int32(4))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = v14 & int32(2147483647)
	if base.Ui32(v19) < base.Ui32(int32(101)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = int32(0)
	if v14 < v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v27 = v19
	goto L8
L7:
	;
	v27 = v14 << (uint(int32(1)) % 32)
	goto L8
L8:
	;
	v31 = v27<<(uint(int32(3))%32) + int32(8)
	v32 = F_palloc(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31 << (uint(int32(2)) % 32)
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v22
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v10 + int32(16)
	return v32
L13:
	;
	v50 = F_pq_getmsgfloat8(m, v12)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32+int32(8)+v40<<(uint(int32(3))%32)))) = v50
	v54 = v40 + int32(1)
	if v54 != v27 {
		v40 = v54
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg(m, int32(398092), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(100)
	F_errdetail(m, int32(573688), v10)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(496403), int32(371), int32(36564))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_scanner_finish(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_cube_yylex_destroy(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_cube_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v63 float64
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	F_enlargeStringInfo(m, v8, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v25 = int32(24)
	v27 = int32(65280)
	v29 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v22+v23))) = v18<<(uint(v25)%32) | v18&v27<<(uint(v29)%32) | (int32(base.Ui32(v18)>>(uint(v29)%32))&v27 | int32(base.Ui32(v18)>>(uint(v25)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v22 + int32(4)
	v45 = v15 & int32(2147483647)
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v55 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 << (uint(int32(2)) % 32)
	goto L12
L8:
	;
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v11+int32(8)+v55<<(uint(int32(3))%32))))
	F_pq_sendfloat8(m, v8, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v67 = v55 + int32(1)
	if v67 != v45<<(uint(int32(base.Ui32(v48^int32(-1))>>(uint(int32(31))%32)))%32) {
		v55 = v67
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	m.G0 = v8 + int32(16)
	return v75
}
func F_cube_union_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 float64
	_ = v103
	var v105 int32
	_ = v105
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v137 float64
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v157 float64
	_ = v157
	var v159 int32
	_ = v159
	var v164 float64
	_ = v164
	var v169 float64
	_ = v169
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v180 float64
	_ = v180
	var v186 int32
	_ = v186
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 float64
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 float64
	_ = v207
	var v210 int32
	_ = v210
	var v218 float64
	_ = v218
	var v221 int32
	_ = v221
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 float64
	_ = v283
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 float64
	_ = v300
	var v302 int32
	_ = v302
	var v303 float64
	_ = v303
	var v306 float64
	_ = v306
	var v308 float64
	_ = v308
	var v310 float64
	_ = v310
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v316 float64
	_ = v316
	var v319 float64
	_ = v319
	var v321 float64
	_ = v321
	var v324 float64
	_ = v324
	var v326 float64
	_ = v326
	var v327 float64
	_ = v327
	var v330 float64
	_ = v330
	var v334 float64
	_ = v334
	var v337 int32
	_ = v337
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v396 int32
	_ = v396
	var v399 float64
	_ = v399
	var v404 float64
	_ = v404
	var v407 int32
	_ = v407
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = int32(2147483647)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = base.B2i32(base.Ui32(v28&v29) < base.Ui32(v31&v29))
	if base.Ui32(v28&v29) < base.Ui32(v31&v29) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = l1
	goto L6
L5:
	;
	v35 = l0
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v36<<(uint(int32(4))%32) | int32(8)
	v41 = F_palloc0(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(int32(2)) % 32)
	v49 = v36 & int32(2147483647)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v53 = v49 | v50&int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if base.Ui32(v28&v29) < base.Ui32(v31&v29) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.Ui32(v59) < base.Ui32(v236) {
		goto L62
	} else {
		goto L63
	}
L10:
	;
	v56 = l0
	goto L12
L11:
	;
	v56 = l1
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v59 = v57 & int32(2147483647)
	if v59 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v236 = v55 & int32(2147483647)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v65 = v55 & int32(2147483647)
	v66 = int32(8)
	v67 = v41 + v66
	v69 = v56 + v66
	v71 = v35 + v66
	v72 = int32(0)
	v75 = v72
	goto L16
L16:
	;
	v101 = v75 << (uint(int32(3)) % 32)
	v102 = v71 + v101
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v102)))
	v105 = base.B2i32(v55 < int32(0))
	if v55 < int32(0) {
		v113 = v103
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v236 = v65
	goto L9
L18:
	;
	v114 = v101 + v69
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	v116 = int32(0)
	v117 = base.B2i32(v57 < v116)
	if v117 == v116 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v71+(v75+v55)<<(uint(int32(3))%32))))
	if base.F64_gt(v110, v103) != 0 {
		v113 = v103
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v113 = v110
	goto L18
L21:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v114)))
	if v117 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L22:
	;
	v175 = v169
	v176 = v169
	goto L21
L23:
	;
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v71+v159<<(uint(int32(3))%32))))
	if base.F64_gt(v157, v164) == int32(0) {
		v175 = v157
		v176 = v164
		goto L21
	} else {
		goto L42
	}
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v101+v67))) = v149
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v102)))
	if v55 < int32(0) {
		v169 = v154
		goto L22
	} else {
		goto L41
	}
L25:
	;
	if base.B2i32(v72 <= v55) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v124 = *(*float64)(unsafe.Add(mBase, uint32(v69+(v75+v57)<<(uint(int32(3))%32))))
	v125 = base.F64_lt(v115, v124)
	if v125 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if base.F64_gt(v115, v113) == int32(0) {
		v149 = v115
		goto L24
	} else {
		goto L34
	}
L29:
	;
	v126 = v115
	goto L31
L30:
	;
	v126 = v124
	goto L31
L31:
	;
	if base.F64_lt(v113, v126) != 0 {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	if v125 != 0 {
		v149 = v115
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v149 = v124
	goto L24
L34:
	;
	goto L25
L35:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v101+v67))) = v103
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v102)))
	v175 = v137
	v176 = v137
	goto L21
L36:
	;
	goto L37
L37:
	;
	v139 = v75 + v55
	v142 = v71 + v139<<(uint(int32(3))%32)
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v142)))
	if base.F64_lt(v103, v143) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v145 = v102
	goto L40
L39:
	;
	v145 = v142
	goto L40
L40:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v145)))
	*(*float64)(unsafe.Add(mBase, uint32(v101+v67))) = v146
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v102)))
	v157 = v148
	v159 = v139
	goto L23
L41:
	;
	v157 = v154
	v159 = v75 + v55
	goto L23
L42:
	;
	v169 = v157
	goto L22
L43:
	;
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v210)))
	*(*float64)(unsafe.Add(mBase, uint32(v67+(v75+v65)<<(uint(int32(3))%32)))) = v218
	v221 = v75 + int32(1)
	if v221 != v59 {
		v75 = v221
		goto L16
	} else {
		goto L61
	}
L44:
	;
	if v188 != 0 {
		v210 = v114
		goto L43
	} else {
		goto L60
	}
L45:
	;
	v206 = v71 + (v75+v55)<<(uint(int32(3))%32)
	v207 = *(*float64)(unsafe.Add(mBase, uint32(v206)))
	if base.F64_gt(v175, v207) != 0 {
		v210 = v102
		goto L43
	} else {
		goto L59
	}
L46:
	;
	v186 = v69 + (v75+v57)<<(uint(int32(3))%32)
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v186)))
	v188 = base.F64_gt(v180, v187)
	if v188 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v195 = base.F64_lt(v180, v176)
	if v195 != 0 {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v189 = v180
	goto L51
L50:
	;
	v189 = v187
	goto L51
L51:
	;
	if base.F64_gt(v176, v189) == int32(0) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	if v105 == int32(0) {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v210 = v102
	goto L43
L54:
	;
	v196 = v102
	goto L56
L55:
	;
	v196 = v114
	goto L56
L56:
	;
	if v55 < int32(0) {
		v210 = v196
		goto L43
	} else {
		goto L57
	}
L57:
	;
	if v195 == int32(0) {
		v210 = v196
		goto L43
	} else {
		goto L58
	}
L58:
	;
	goto L45
L59:
	;
	v210 = v206
	goto L43
L60:
	;
	v210 = v186
	goto L43
L61:
	;
	goto L17
L62:
	;
	v249 = int32(8)
	v250 = v41 + v249
	v252 = v35 + v249
	v256 = v59
	goto L65
L63:
	;
	goto L64
L64:
	;
	v364 = int32(0)
	if v50 < v364 {
		goto L94
	} else {
		goto L95
	}
L65:
	;
	v280 = int32(3)
	v281 = v256 << (uint(v280) % 32)
	v282 = v252 + v281
	v283 = *(*float64)(unsafe.Add(mBase, uint32(v282)))
	if base.B2i32(v55 < int32(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L64
L67:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v250+(v256+v49)<<(uint(v280)%32)))) = v334
	v337 = v256 + int32(1)
	if v337 != v236 {
		v256 = v337
		goto L65
	} else {
		goto L92
	}
L68:
	;
	v293 = int32(3)
	v295 = v252 + (v256+v55)<<(uint(v293)%32)
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v252+(v256+v236)<<(uint(v293)%32))))
	if base.F64_lt(v283, v300) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v321 = float64(0)
	if base.F64_gt(v283, v321) != 0 {
		goto L86
	} else {
		goto L87
	}
L71:
	;
	v302 = v282
	goto L73
L72:
	;
	v302 = v295
	goto L73
L73:
	;
	v303 = *(*float64)(unsafe.Add(mBase, uint32(v302)))
	if base.F64_gt(v303, float64(0)) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v310 = float64(0)
	goto L76
L75:
	;
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v295)))
	if base.F64_lt(v283, v306) != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v281+v250))) = v310
	v313 = *(*float64)(unsafe.Add(mBase, uint32(v282)))
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v295)))
	if base.F64_gt(v313, v314) != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v308 = v283
	goto L79
L78:
	;
	v308 = v306
	goto L79
L79:
	;
	v310 = v308
	goto L76
L80:
	;
	v316 = v313
	goto L82
L81:
	;
	v316 = v314
	goto L82
L82:
	;
	if base.F64_lt(v316, float64(0)) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v319 = float64(0)
	goto L85
L84:
	;
	v319 = v316
	goto L85
L85:
	;
	v334 = v319
	goto L67
L86:
	;
	v324 = v321
	goto L88
L87:
	;
	v324 = v283
	goto L88
L88:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v281+v250))) = v324
	v326 = float64(0)
	v327 = *(*float64)(unsafe.Add(mBase, uint32(v282)))
	if base.F64_lt(v327, v326) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v330 = v326
	goto L91
L90:
	;
	v330 = v327
	goto L91
L91:
	;
	v334 = v330
	goto L67
L92:
	;
	goto L66
L93:
	;
	return v41
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v36 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36<<(uint(int32(5))%32) + int32(32)
	goto L93
L95:
	;
	if v53 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v370 = v41 + int32(8)
	v371 = v364
	goto L97
L97:
	;
	v396 = int32(3)
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v370+v371<<(uint(v396)%32))))
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v370+(v371+v53)<<(uint(v396)%32))))
	if base.F64_ne(v399, v404) != 0 {
		goto L93
	} else {
		goto L99
	}
L98:
	;
	goto L94
L99:
	;
	v407 = v371 + int32(1)
	if v407 != v53 {
		v371 = v407
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
}
func F_cube_yy_scan_bytes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	v4 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(-2)) {
		v14 = l1 + int32(2)
		v15 = F_palloc(m, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				F_yy_fatal_error_6(m, int32(664961))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if l1 == int32(0) {
				} else {
					if base.Ui32(int32(4)) <= base.Ui32(l1) {
						v30 = v4
						v34 = v4
						for {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v30))))
							*(*uint8)(unsafe.Add(mBase, uint32(v30+v15))) = uint8(v39)
							v42 = v30 | int32(1)
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v42))))
							*(*uint8)(unsafe.Add(mBase, uint32(v15+v42))) = uint8(v45)
							v48 = v30 | int32(2)
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v48))))
							*(*uint8)(unsafe.Add(mBase, uint32(v15+v48))) = uint8(v51)
							v54 = v30 | int32(3)
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v54))))
							*(*uint8)(unsafe.Add(mBase, uint32(v15+v54))) = uint8(v57)
							v59 = int32(4)
							v60 = v30 + v59
							v62 = v34 + v59
							if v62 != l1&int32(-4) {
								v30 = v60
								v34 = v62
								continue
							} else {
								break
							}
							break
						}
						v67 = v60
					} else {
						v67 = v4
					}
					v75 = l1 & int32(3)
					if v75 == int32(0) {
					} else {
						v81 = v67
						v84 = v4
						for {
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v81))))
							*(*uint8)(unsafe.Add(mBase, uint32(v81+v15))) = uint8(v90)
							v92 = int32(1)
							v95 = v84 + v92
							if v95 != v75 {
								v81 = v81 + v92
								v84 = v95
								continue
							} else {
								break
							}
							break
						}
					}
				}
				v108 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1+v15))) = uint16(v108)
				v110 = F_cube_yy_scan_buffer(m, v15, v14, l2)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					if v110 == int32(0) {
						F_yy_fatal_error_6(m, int32(665002))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v110)+20)) = int32(1)
						return v110
					}
				}
			}
		}
	} else {
		F_yy_fatal_error_6(m, int32(665032))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_cube_yy_scan_string(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = F_strlen(m, l0)
	v4 = F_cube_yy_scan_bytes(m, l0, v3, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_cube_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_6(m, int32(665681))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_6(m, int32(665681))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_cube_yyget_lval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	return v2
}
func F_cube_yyget_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	return v2
}
func F_cube_yyset_debug(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = l0
	return
}
func F_cube_yyset_in(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	return
}
