package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_contains(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v57 float64
	_ = v57
	var v66 float64
	_ = v66
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v120 float64
	_ = v120
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
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
	v13 = int32(0)
	if v6 == v13 {
		v163 = v13
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v188 != v6 {
		goto L42
	} else {
		goto L43
	}
L5:
	;
	v187 = v163
	goto L4
L6:
	;
	if v11 == int32(0) {
		v163 = v13
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v11 + int32(8)
	v42 = v31
	goto L11
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v37+v42<<(uint(int32(3))%32))))
	if base.F64_ne(v57, float64(0)) != 0 {
		v163 = v13
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v37+(v42+v34)<<(uint(int32(3))%32))))
	if base.F64_ne(v66, float64(0)) != 0 {
		v163 = v13
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v70 = v42 + int32(1)
	if v70 != v34 {
		v42 = v70
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	goto L12
L19:
	;
	v87 = v31
	goto L21
L20:
	;
	v87 = v34
	goto L21
L21:
	;
	if v87 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v187 = int32(1)
	goto L4
L23:
	;
	goto L24
L24:
	;
	v91 = int32(8)
	v92 = v11 + v91
	v94 = v6 + v91
	v98 = int32(0)
	goto L25
L25:
	;
	v111 = v98 << (uint(int32(3)) % 32)
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v94+v111)))
	v115 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v123 = v113
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v163 = v155
	goto L5
L27:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v92+v111)))
	v127 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v135 = v125
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v94+(v98+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v113, v120) != 0 {
		v123 = v113
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v123 = v120
	goto L27
L30:
	;
	v136 = int32(0)
	if base.F64_gt(v123, v135) != 0 {
		v163 = v136
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v92+(v98+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v125, v132) != 0 {
		v135 = v125
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v135 = v132
	goto L30
L33:
	;
	if v29 < int32(0) {
		v144 = v113
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v32 < int32(0) {
		v152 = v125
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v94+(v98+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v113, v142) != 0 {
		v144 = v113
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v144 = v142
	goto L34
L37:
	;
	if base.F64_gt(v152, v144) != 0 {
		v163 = v136
		goto L5
	} else {
		goto L40
	}
L38:
	;
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v92+(v98+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v125, v150) != 0 {
		v152 = v125
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v152 = v150
	goto L37
L40:
	;
	v155 = int32(1)
	v157 = v98 + v155
	if v157 != v87 {
		v98 = v157
		goto L25
	} else {
		goto L41
	}
L41:
	;
	goto L26
L42:
	;
	F_pfree(m, v6)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v192 != v11 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	F_pfree(m, v11)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	return v187
L49:
	;
	goto L48
}
func F_cube_gt(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(int32(0) < v487)
L130:
	;
	goto L129
}
func F_cube_is_point(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v34 float64
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v7 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v46 != v9 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v13 < int32(0) {
		v42 = v7
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v13 == int32(0) {
		v42 = v7
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = v9 + int32(8)
	v21 = int32(0)
	goto L6
L6:
	;
	v26 = int32(3)
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v19+v21<<(uint(v26)%32))))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v19+(v21+v13)<<(uint(v26)%32))))
	if base.F64_eq(v29, v34) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v42 = int32(0)
	goto L1
L8:
	;
	v37 = v21 + int32(1)
	if v13 != v37 {
		v21 = v37
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v42 = v7
	goto L1
L12:
	;
	F_pfree(m, v9)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	return v42
L15:
	;
	goto L14
}
func F_cube_ne(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(v487 != int32(0))
L130:
	;
	goto L129
}
func F_cube_overlap(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v65 int32
	_ = v65
	var v67 float64
	_ = v67
	var v69 int32
	_ = v69
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v81 int32
	_ = v81
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v151 float64
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 float64
	_ = v162
	var v164 int32
	_ = v164
	var v165 float64
	_ = v165
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v173 float64
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
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
	v17 = int32(0)
	if v6 == v17 {
		v193 = v17
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v215 != v6 {
		goto L55
	} else {
		goto L56
	}
L5:
	;
	v214 = v193
	goto L4
L6:
	;
	if v11 == int32(0) {
		v193 = v17
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v35 = base.B2i32(base.Ui32(v29&v30) < base.Ui32(v32&v30))
	if base.Ui32(v29&v30) < base.Ui32(v32&v30) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v36 = v11
	goto L10
L9:
	;
	v36 = v6
	goto L10
L10:
	;
	if base.Ui32(v29&v30) < base.Ui32(v32&v30) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v127 = v118 & int32(2147483647)
	if base.Ui32(v127) <= base.Ui32(v40) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v37 = v6
	goto L14
L13:
	;
	v37 = v11
	goto L14
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v40 = v38 & int32(2147483647)
	if v40 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v118 = v43
	goto L11
L16:
	;
	goto L17
L17:
	;
	v44 = int32(8)
	v45 = v37 + v44
	v47 = v36 + v44
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v50 = int32(0)
	goto L18
L18:
	;
	v65 = v50 << (uint(int32(3)) % 32)
	v67 = *(*float64)(unsafe.Add(mBase, uint32(v47+v65)))
	v69 = base.B2i32(v48 < int32(0))
	if v48 < int32(0) {
		v77 = v67
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v118 = v48
	goto L11
L20:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v65+v45)))
	v81 = base.B2i32(v38 < int32(0))
	if v38 < int32(0) {
		v89 = v79
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v47+(v50+v48)<<(uint(int32(3))%32))))
	if base.F64_lt(v67, v74) != 0 {
		v77 = v67
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v77 = v74
	goto L20
L23:
	;
	v90 = int32(0)
	if base.F64_gt(v77, v89) != 0 {
		v193 = v90
		goto L5
	} else {
		goto L26
	}
L24:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v45+(v50+v38)<<(uint(int32(3))%32))))
	if base.F64_gt(v79, v86) != 0 {
		v89 = v79
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = v86
	goto L23
L26:
	;
	if v48 < int32(0) {
		v98 = v67
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v38 < int32(0) {
		v106 = v79
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v47+(v50+v48)<<(uint(int32(3))%32))))
	if base.F64_gt(v67, v96) != 0 {
		v98 = v67
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v98 = v96
	goto L27
L30:
	;
	if base.F64_lt(v98, v106) != 0 {
		v193 = v90
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v45+(v50+v38)<<(uint(int32(3))%32))))
	if base.F64_lt(v79, v104) != 0 {
		v106 = v79
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v106 = v104
	goto L30
L33:
	;
	v110 = v50 + int32(1)
	if v110 != v40 {
		v50 = v110
		goto L18
	} else {
		goto L34
	}
L34:
	;
	goto L19
L35:
	;
	v214 = int32(1)
	goto L4
L36:
	;
	goto L37
L37:
	;
	v131 = v36 + int32(8)
	v135 = v40
	goto L38
L38:
	;
	v150 = v131 + v135<<(uint(int32(3))%32)
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v150)))
	if base.B2i32(v118 < int32(0)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v193 = int32(0)
	goto L5
L40:
	;
	goto L39
L41:
	;
	if base.F64_lt(v173, float64(0)) != 0 {
		goto L40
	} else {
		goto L53
	}
L42:
	;
	v155 = int32(3)
	v157 = v131 + (v135+v118)<<(uint(v155)%32)
	v162 = *(*float64)(unsafe.Add(mBase, uint32(v131+(v135+v127)<<(uint(v155)%32))))
	if base.F64_lt(v151, v162) != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if base.F64_gt(v151, float64(0)) != 0 {
		goto L40
	} else {
		goto L52
	}
L45:
	;
	v164 = v150
	goto L47
L46:
	;
	v164 = v157
	goto L47
L47:
	;
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v164)))
	if base.F64_gt(v165, float64(0)) != 0 {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v157)))
	if base.F64_gt(v151, v168) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v170 = v151
	goto L51
L50:
	;
	v170 = v168
	goto L51
L51:
	;
	v173 = v170
	goto L41
L52:
	;
	v173 = v151
	goto L41
L53:
	;
	v178 = int32(1)
	v180 = v135 + v178
	if v127 != v180 {
		v135 = v180
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v193 = v178
	goto L5
L55:
	;
	F_pfree(m, v6)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v219 != v11 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	F_pfree(m, v11)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	return v214
L62:
	;
	goto L61
}
func F_cube_yy_delete_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v4 == int32(0) {
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v10 = v4 + v7<<(uint(int32(2))%32)
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if l0 != v11 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
			}
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16 == int32(0) {
			F_pfree(m, l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v19 == int32(0) {
				F_pfree(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					return
				}
			} else {
				F_pfree(m, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_cube_yy_scan_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v4 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v77 = v4
		return v77
	} else {
		v9 = l1 - int32(2)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v9))))
		if v11 != 0 {
			v77 = v4
			return v77
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1-int32(1)))))
			if v15 != 0 {
				v77 = v4
				return v77
			} else {
				v17 = F_palloc(m, int32(48))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						F_yy_fatal_error_6(m, int32(683554))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v23 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v23
						*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v9
						*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = v23
						F_cube_yyensure_buffer_stack(m, l2)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(int32(2))%32))))
							if v42 == v17 {
								v77 = v17
							} else {
								if v42 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)))
									*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v45)
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48<<(uint(int32(2))%32))))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v53
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v55
									v57 = v48
									v58 = v47
								} else {
									v57 = v38
									v58 = v37
								}
								*(*int32)(unsafe.Add(mBase, uint32(v58+v57<<(uint(int32(2))%32)))) = v17
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v64
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v66
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v69
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)) = uint8(v71)
								v77 = v17
							}
							return v77
						}
					}
				}
			}
		}
	}
}
func F_cube_yyget_column(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == v2 {
		v16 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4+v7<<(uint(int32(2))%32))))
		if v11 == int32(0) {
			v16 = v2
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
			v16 = v14
		}
	}
	return v16
}
func F_cube_yyget_leng(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	return v2
}
func F_cube_yyparse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1000 int32
	_ = v1000
	var v1008 int32
	_ = v1008
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1037 int32
	_ = v1037
	v23 = m.G0
	v25 = v23 - int32(1312)
	m.G0 = v25
	v31 = v25 + int32(96)
	v33 = v25 + int32(896)
	v39 = int32(0)
	v40 = v33
	v42 = v31
	v48 = v33
	v49 = int32(-2)
	v50 = int32(200)
	v54 = v31
	goto L1
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v39)
	v59 = v50 << (uint(int32(1)) % 32)
	if base.Ui32(v48+v59-int32(2)) <= base.Ui32(v40) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	if v25+int32(896) != v1024 {
		goto L237
	} else {
		goto L238
	}
L3:
	;
	goto L2
L4:
	;
	v39 = v1008
	v40 = v991 + int32(2)
	v42 = v993
	v48 = v116
	v49 = v1000
	v50 = v117
	v54 = v118
	goto L1
L5:
	;
	v954 = v113 - v182<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v954)+4)) = v946
	v957 = v954 + int32(4)
	v960 = v111 - v182<<(uint(int32(1))%32)
	v961 = int32(*(*int16)(unsafe.Add(mBase, uint32(v960))))
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+uint32(_consts[1320]))))
	v970 = v968 - int32(9)
	v972 = int32(*(*int8)(unsafe.Add(mBase, uint32(v970)+uint32(_consts[1321]))))
	v973 = v961 + v972
	if base.Ui32(int32(17)) < base.Ui32(v973) {
		goto L234
	} else {
		goto L235
	}
L6:
	;
	v927 = F_write_point_as_box(m, v909, v614, l0, l2)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L16
	} else {
		goto L232
	}
L7:
	;
	v902 = F_write_point_as_box(m, v884, v695, l0, l2)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L16
	} else {
		goto L230
	}
L8:
	;
	F_cube_yyerror(m, l0, l1, l2, l3, int32(440744))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L16
	} else {
		goto L229
	}
L9:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v50) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v111 = v40
	v113 = v42
	v116 = v48
	v117 = v50
	v118 = v54
	goto L11
L11:
	;
	v121 = int32(1) << (uint(v39) % 32)
	if v121&int32(369794) != 0 {
		v169 = v49
		goto L34
	} else {
		goto L35
	}
L12:
	;
	v66 = int32(10000)
	if base.Ui32(v66) <= base.Ui32(v59) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v69 = v66
	goto L15
L14:
	;
	v69 = v59
	goto L15
L15:
	;
	v74 = F_palloc(m, v69*int32(6)|int32(3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	if v74 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v81 = int32(1)
	v84 = (v40-v48)>>(uint(v81)%32) + v81
	v86 = v84 << (uint(v81) % 32)
	if v86 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v91 = v88 + v69<<(uint(int32(1))%32)
	v93 = v84 << (uint(int32(2)) % 32)
	if v93 != 0 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v87 = F__emscripten_memcpy_bulkmem(m, v74, v48, v86)
	mBase = m.M
	v88 = v87
	goto L22
L21:
	;
	v88 = v74
	goto L22
L22:
	;
	goto L19
L23:
	;
	if v25+int32(896) != v48 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v94 = F__emscripten_memcpy_bulkmem(m, v91, v54, v93)
	mBase = m.M
	v95 = v94
	goto L26
L25:
	;
	v95 = v91
	goto L26
L26:
	;
	goto L23
L27:
	;
	F_pfree(m, v48)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v69) <= base.Ui32(v84) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v1019 = int32(1)
	v1024 = v88
	goto L3
L32:
	;
	goto L33
L33:
	;
	v111 = v88 + v84<<(uint(int32(1))%32) - int32(2)
	v113 = v93 + v95 - int32(4)
	v116 = v88
	v117 = v69
	v118 = v95
	goto L11
L34:
	;
	if v121&int32(154397) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L35:
	;
	v127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1322]))))
	if v49 == int32(-2) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v147 = v127 + v146
	if base.Ui32(int32(17)) < base.Ui32(v147) {
		v169 = v145
		goto L34
	} else {
		goto L45
	}
L37:
	;
	v132 = F_cube_yylex(m, v25+int32(1308), l3)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L16
	} else {
		goto L40
	}
L38:
	;
	v134 = v49
	goto L39
L39:
	;
	if v134 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v134 = v132
	goto L39
L41:
	;
	v137 = int32(0)
	v145 = v137
	v146 = v137
	goto L36
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(263)) < base.Ui32(v134) {
		v145 = v134
		v146 = int32(2)
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[1323]))))
	v145 = v134
	v146 = v144
	goto L36
L45:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[1324]))))
	if v146 != v152 {
		v169 = v145
		goto L34
	} else {
		goto L46
	}
L46:
	;
	if v147 != int32(10) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1308))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v156
	if v145 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v1019 = int32(0)
	v1024 = v116
	goto L3
L50:
	;
	v160 = int32(-2)
	goto L52
L51:
	;
	v160 = int32(0)
	goto L52
L52:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[1325]))))
	v991 = v111
	v993 = v113 + int32(4)
	v1000 = v160
	v1008 = v165
	goto L4
L53:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[1326]))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+uint32(_consts[1327]))))
	v184 = int32(2)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v113+(int32(1)-v182)<<(uint(v184)%32))))
	switch v179 - v184 {
	case 0:
		goto L63
	case 1:
		goto L62
	case 2:
		goto L61
	case 3:
		goto L60
	case 4:
		goto L59
	case 5:
		goto L58
	case 6:
		goto L57
	case 7:
		goto L56
	default:
		v946 = v187
		goto L5
	}
L54:
	;
	goto L55
L55:
	;
	F_cube_yyerror(m, l0, l1, l2, l3, int32(212048))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L16
	} else {
		goto L228
	}
L56:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v113-int32(8))))
	v862 = F_strlen(m, v861)
	mBase = m.M
	v864 = int32(44)
	*(*uint16)(unsafe.Add(mBase, uint32(v862+v861))) = uint16(v864)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v867 = F_strlen(m, v861)
	mBase = m.M
	v869 = F_strcpy(m, v867+v861, v866)
	mBase = m.M
	goto L227
L57:
	;
	v782 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L16
	} else {
		goto L205
	}
L58:
	;
	v780 = F_pstrdup(m, int32(757108))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L16
	} else {
		goto L204
	}
L59:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v113-int32(4))))
	v946 = v778
	goto L5
L60:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
	if v696 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L61:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if v615 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L62:
	;
	v402 = int32(0)
	v405 = v113 - int32(8)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v407 == v402 {
		v459 = v402
		goto L112
	} else {
		goto L113
	}
L63:
	;
	v190 = int32(0)
	v193 = v113 - int32(12)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v195 == v190 {
		v247 = v190
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v266 = v113 - int32(4)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v268 == int32(0) {
		v322 = v190
		goto L78
	} else {
		goto L79
	}
L65:
	;
	v198 = int32(1)
	v199 = int32(44)
	v200 = F___strchrnul(m, v194, v199)
	mBase = m.M
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v202 == v199 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v206 == int32(0) {
		v247 = v198
		goto L64
	} else {
		goto L70
	}
L67:
	;
	v206 = v200
	goto L69
L68:
	;
	v206 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v213 = v198
	v219 = v206
	goto L71
L71:
	;
	v231 = int32(1)
	v232 = v213 + v231
	v235 = int32(44)
	v236 = F___strchrnul(m, v219+v231, v235)
	mBase = m.M
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if v238 == v235 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v247 = v232
	goto L64
L73:
	;
	if v242 != 0 {
		v213 = v232
		v219 = v242
		goto L71
	} else {
		goto L77
	}
L74:
	;
	v242 = v236
	goto L76
L75:
	;
	v242 = int32(0)
	goto L76
L76:
	;
	goto L73
L77:
	;
	goto L72
L78:
	;
	if v247 != v322 {
		goto L92
	} else {
		goto L93
	}
L79:
	;
	v271 = int32(1)
	v272 = int32(44)
	v273 = F___strchrnul(m, v267, v272)
	mBase = m.M
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v275 == v272 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v279 == int32(0) {
		v322 = v271
		goto L78
	} else {
		goto L84
	}
L81:
	;
	v279 = v273
	goto L83
L82:
	;
	v279 = int32(0)
	goto L83
L83:
	;
	goto L80
L84:
	;
	v288 = v271
	v292 = v279
	goto L85
L85:
	;
	v304 = int32(1)
	v305 = v288 + v304
	v308 = int32(44)
	v309 = F___strchrnul(m, v292+v304, v308)
	mBase = m.M
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v311 == v308 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v322 = v305
	goto L78
L87:
	;
	if v315 != 0 {
		v288 = v305
		v292 = v315
		goto L85
	} else {
		goto L91
	}
L88:
	;
	v315 = v309
	goto L90
L89:
	;
	v315 = int32(0)
	goto L90
L90:
	;
	goto L87
L91:
	;
	goto L86
L92:
	;
	v339 = int32(1)
	v340 = F_errsave_start(m, l2)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L16
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if int32(101) <= v247 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	if v340 == int32(0) {
		v1019 = v339
		v1024 = v116
		goto L3
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L16
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(419694), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v353
	F_errdetail(m, int32(660997), v25+int32(16))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	F_errsave_finish(m, l2, int32(27008), int32(58), int32(360688))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	v1019 = v339
	v1024 = v116
	goto L3
L101:
	;
	v372 = int32(1)
	v373 = F_errsave_start(m, l2)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L16
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v399 = F_write_box(m, v247, v194, v267, l0, l2)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L16
	} else {
		goto L110
	}
L104:
	;
	if v373 == int32(0) {
		v1019 = v372
		v1024 = v116
		goto L3
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(419694), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(100)
	F_errdetail(m, int32(591865), v25)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	F_errsave_finish(m, l2, int32(27008), int32(67), int32(360688))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	v1019 = v372
	v1024 = v116
	goto L3
L110:
	;
	if v399 != 0 {
		v946 = v187
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v1019 = int32(1)
	v1024 = v116
	goto L3
L112:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v478 == int32(0) {
		v532 = v402
		goto L126
	} else {
		goto L127
	}
L113:
	;
	v410 = int32(1)
	v411 = int32(44)
	v412 = F___strchrnul(m, v406, v411)
	mBase = m.M
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	if v414 == v411 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v418 == int32(0) {
		v459 = v410
		goto L112
	} else {
		goto L118
	}
L115:
	;
	v418 = v412
	goto L117
L116:
	;
	v418 = int32(0)
	goto L117
L117:
	;
	goto L114
L118:
	;
	v425 = v410
	v431 = v418
	goto L119
L119:
	;
	v443 = int32(1)
	v444 = v425 + v443
	v447 = int32(44)
	v448 = F___strchrnul(m, v431+v443, v447)
	mBase = m.M
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	if v450 == v447 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v459 = v444
	goto L112
L121:
	;
	if v454 != 0 {
		v425 = v444
		v431 = v454
		goto L119
	} else {
		goto L125
	}
L122:
	;
	v454 = v448
	goto L124
L123:
	;
	v454 = int32(0)
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L120
L126:
	;
	if v459 != v532 {
		goto L140
	} else {
		goto L141
	}
L127:
	;
	v481 = int32(1)
	v482 = int32(44)
	v483 = F___strchrnul(m, v477, v482)
	mBase = m.M
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	if v485 == v482 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if v489 == int32(0) {
		v532 = v481
		goto L126
	} else {
		goto L132
	}
L129:
	;
	v489 = v483
	goto L131
L130:
	;
	v489 = int32(0)
	goto L131
L131:
	;
	goto L128
L132:
	;
	v498 = v481
	v502 = v489
	goto L133
L133:
	;
	v514 = int32(1)
	v515 = v498 + v514
	v518 = int32(44)
	v519 = F___strchrnul(m, v502+v514, v518)
	mBase = m.M
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if v521 == v518 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v532 = v515
	goto L126
L135:
	;
	if v525 != 0 {
		v498 = v515
		v502 = v525
		goto L133
	} else {
		goto L139
	}
L136:
	;
	v525 = v519
	goto L138
L137:
	;
	v525 = int32(0)
	goto L138
L138:
	;
	goto L135
L139:
	;
	goto L134
L140:
	;
	v549 = int32(1)
	v550 = F_errsave_start(m, l2)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L16
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if int32(101) <= v459 {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	if v550 == int32(0) {
		v1019 = v549
		v1024 = v116
		goto L3
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L16
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(419694), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v563
	F_errdetail(m, int32(660997), v25+int32(48))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L16
	} else {
		goto L147
	}
L147:
	;
	F_errsave_finish(m, l2, int32(27008), int32(88), int32(360688))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L16
	} else {
		goto L148
	}
L148:
	;
	v1019 = v549
	v1024 = v116
	goto L3
L149:
	;
	v582 = int32(1)
	v583 = F_errsave_start(m, l2)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L16
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v611 = F_write_box(m, v459, v406, v477, l0, l2)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L16
	} else {
		goto L158
	}
L152:
	;
	if v583 == int32(0) {
		v1019 = v582
		v1024 = v116
		goto L3
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L16
	} else {
		goto L154
	}
L154:
	;
	F_errmsg(m, int32(419694), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L16
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(100)
	F_errdetail(m, int32(591865), v25+int32(32))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L16
	} else {
		goto L156
	}
L156:
	;
	F_errsave_finish(m, l2, int32(27008), int32(97), int32(360688))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L16
	} else {
		goto L157
	}
L157:
	;
	v1019 = v582
	v1024 = v116
	goto L3
L158:
	;
	if v611 != 0 {
		v946 = v187
		goto L5
	} else {
		goto L159
	}
L159:
	;
	v1019 = int32(1)
	v1024 = v116
	goto L3
L160:
	;
	v909 = int32(0)
	goto L6
L161:
	;
	goto L162
L162:
	;
	v619 = int32(1)
	v620 = int32(44)
	v621 = F___strchrnul(m, v614, v620)
	mBase = m.M
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	if v623 == v620 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v627 == int32(0) {
		v909 = v619
		goto L6
	} else {
		goto L167
	}
L164:
	;
	v627 = v621
	goto L166
L165:
	;
	v627 = int32(0)
	goto L166
L166:
	;
	goto L163
L167:
	;
	v634 = v619
	v636 = v627
	goto L168
L168:
	;
	v652 = int32(1)
	v653 = v634 + v652
	v656 = int32(44)
	v657 = F___strchrnul(m, v636+v652, v656)
	mBase = m.M
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657))))
	if v659 == v656 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	if base.Ui32(v634) < base.Ui32(int32(100)) {
		v909 = v653
		goto L6
	} else {
		goto L175
	}
L170:
	;
	if v663 != 0 {
		v634 = v653
		v636 = v663
		goto L168
	} else {
		goto L174
	}
L171:
	;
	v663 = v657
	goto L173
L172:
	;
	v663 = int32(0)
	goto L173
L173:
	;
	goto L170
L174:
	;
	goto L169
L175:
	;
	v666 = int32(1)
	v667 = F_errsave_start(m, l2)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L16
	} else {
		goto L176
	}
L176:
	;
	if v667 == int32(0) {
		v1019 = v666
		v1024 = v116
		goto L3
	} else {
		goto L177
	}
L177:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L16
	} else {
		goto L178
	}
L178:
	;
	F_errmsg(m, int32(419694), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L16
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = int32(100)
	F_errdetail(m, int32(591865), v25-int32(-64))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L16
	} else {
		goto L180
	}
L180:
	;
	F_errsave_finish(m, l2, int32(27008), int32(116), int32(360688))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L16
	} else {
		goto L181
	}
L181:
	;
	v1019 = v666
	v1024 = v116
	goto L3
L182:
	;
	v884 = int32(0)
	goto L7
L183:
	;
	goto L184
L184:
	;
	v700 = int32(1)
	v701 = int32(44)
	v702 = F___strchrnul(m, v695, v701)
	mBase = m.M
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702))))
	if v704 == v701 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v708 == int32(0) {
		v884 = v700
		goto L7
	} else {
		goto L189
	}
L186:
	;
	v708 = v702
	goto L188
L187:
	;
	v708 = int32(0)
	goto L188
L188:
	;
	goto L185
L189:
	;
	v715 = v700
	v717 = v708
	goto L190
L190:
	;
	v733 = int32(1)
	v734 = v715 + v733
	v737 = int32(44)
	v738 = F___strchrnul(m, v717+v733, v737)
	mBase = m.M
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738))))
	if v740 == v737 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	if base.Ui32(v715) < base.Ui32(int32(100)) {
		v884 = v734
		goto L7
	} else {
		goto L197
	}
L192:
	;
	if v744 != 0 {
		v715 = v734
		v717 = v744
		goto L190
	} else {
		goto L196
	}
L193:
	;
	v744 = v738
	goto L195
L194:
	;
	v744 = int32(0)
	goto L195
L195:
	;
	goto L192
L196:
	;
	goto L191
L197:
	;
	v747 = int32(1)
	v748 = F_errsave_start(m, l2)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L16
	} else {
		goto L198
	}
L198:
	;
	if v748 == int32(0) {
		v1019 = v747
		v1024 = v116
		goto L3
	} else {
		goto L199
	}
L199:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L16
	} else {
		goto L200
	}
L200:
	;
	F_errmsg(m, int32(419694), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L16
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = int32(100)
	F_errdetail(m, int32(591865), v25+int32(80))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L16
	} else {
		goto L202
	}
L202:
	;
	F_errsave_finish(m, l2, int32(27008), int32(135), int32(360688))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L16
	} else {
		goto L203
	}
L203:
	;
	v1019 = v747
	v1024 = v116
	goto L3
L204:
	;
	v946 = v780
	goto L5
L205:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if (v784^v782)&int32(3) != 0 {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v946 = v782
	goto L5
L207:
	;
	goto L206
L208:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v839))) = uint8(v838)
	if v838&int32(255) == int32(0) {
		goto L207
	} else {
		goto L223
	}
L209:
	;
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784))))
	v837 = v784
	v838 = v790
	v839 = v782
	goto L208
L210:
	;
	goto L211
L211:
	;
	if v784&int32(3) != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v794 = v784
	v796 = v782
	goto L215
L213:
	;
	v808 = v784
	v810 = v782
	goto L214
L214:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v815 = int32(-2139062144)
	if (int32(16843008)-v812|v812)&v815 != v815 {
		v837 = v808
		v838 = v812
		v839 = v810
		goto L208
	} else {
		goto L219
	}
L215:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v797)
	if v797 == int32(0) {
		goto L207
	} else {
		goto L217
	}
L216:
	;
	v808 = v804
	v810 = v802
	goto L214
L217:
	;
	v801 = int32(1)
	v802 = v796 + v801
	v804 = v794 + v801
	if v804&int32(3) != 0 {
		v794 = v804
		v796 = v802
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v820 = v808
	v821 = v812
	v822 = v810
	goto L220
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822))) = v821
	v824 = int32(4)
	v825 = v822 + v824
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v820)+4))
	v828 = v820 + v824
	v832 = int32(-2139062144)
	if (v826|(int32(16843008)-v826))&v832 == v832 {
		v820 = v828
		v821 = v826
		v822 = v825
		goto L220
	} else {
		goto L222
	}
L221:
	;
	v837 = v828
	v838 = v826
	v839 = v825
	goto L208
L222:
	;
	goto L221
L223:
	;
	v846 = v837
	v848 = v839
	goto L224
L224:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v848)+1)) = uint8(v849)
	v851 = int32(1)
	if v849 != 0 {
		v846 = v846 + v851
		v848 = v848 + v851
		goto L224
	} else {
		goto L226
	}
L225:
	;
	goto L207
L226:
	;
	goto L225
L227:
	;
	v946 = v861
	goto L5
L228:
	;
	v1019 = int32(1)
	v1024 = v116
	goto L3
L229:
	;
	v1019 = int32(2)
	v1024 = v48
	goto L3
L230:
	;
	if v902 != 0 {
		v946 = v187
		goto L5
	} else {
		goto L231
	}
L231:
	;
	v1019 = int32(1)
	v1024 = v116
	goto L3
L232:
	;
	if v927 != 0 {
		v946 = v187
		goto L5
	} else {
		goto L233
	}
L233:
	;
	v1019 = int32(1)
	v1024 = v116
	goto L3
L234:
	;
	v985 = int32(*(*int8)(unsafe.Add(mBase, uint32(v970)+uint32(_consts[1328]))))
	v991 = v960
	v993 = v957
	v1000 = v169
	v1008 = v985
	goto L4
L235:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_consts[1324]))))
	if v978 != v961 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973)+uint32(_consts[1325]))))
	v991 = v960
	v993 = v957
	v1000 = v169
	v1008 = v982
	goto L4
L237:
	;
	F_pfree(m, v1024)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L16
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	m.G0 = v25 + int32(1312)
	return v1019
L240:
	;
	goto L239
}
func F_cube_yypush_buffer_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	if l0 != 0 {
		F_cube_yyensure_buffer_stack(m, l1)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
			if v13 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v15)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18<<(uint(int32(2))%32))))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v25
				v28 = v18 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v28
				v30 = v17
				v31 = v28
			} else {
				v30 = v8
				v31 = v9
			}
			*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = l0
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v44)
			return
		}
	} else {
		return
	}
}
func F_cube_yyrestart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
		if v13 != 0 {
			v27 = v13
			v29 = *(*int32)(unsafe.Add(mBase, _consts[40]))
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
			v40 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
			if v50 == v27 {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				v70 = v40
				v72 = int32(40)
			} else {
				v63 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
				v70 = int32(0)
				v72 = int32(36)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
			*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[40])) = v29
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
			return
		} else {
			F_cube_yyensure_buffer_stack(m, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v19 = F_cube_yy_create_buffer(m, v17, int32(16384), l1)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = v19
					v27 = v19
					v29 = *(*int32)(unsafe.Add(mBase, _consts[40]))
					v30 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
					if v50 == v27 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
						v70 = v40
						v72 = int32(40)
					} else {
						v63 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
						v70 = int32(0)
						v72 = int32(36)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
					*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[40])) = v29
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
					return
				}
			}
		}
	} else {
		F_cube_yyensure_buffer_stack(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v19 = F_cube_yy_create_buffer(m, v17, int32(16384), l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = v19
				v27 = v19
				v29 = *(*int32)(unsafe.Add(mBase, _consts[40]))
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
				v40 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
				if v50 == v27 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
					v70 = v40
					v72 = int32(40)
				} else {
					v63 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
					v70 = int32(0)
					v72 = int32(36)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[40])) = v29
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
				return
			}
		}
	}
}
func F_cube_yyset_column(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32))))
		if v9 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l0
			return
		} else {
			F_yy_fatal_error_6(m, int32(225629))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_yy_fatal_error_6(m, int32(225629))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_cube_yyset_out(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l0
	return
}
